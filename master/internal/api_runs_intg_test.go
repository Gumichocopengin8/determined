//go:build integration
// +build integration

package internal

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/determined-ai/determined/master/internal/db"
	"github.com/determined-ai/determined/master/pkg/model"
	"github.com/determined-ai/determined/master/pkg/ptrs"
	"github.com/determined-ai/determined/master/pkg/schemas"
	"github.com/determined-ai/determined/master/pkg/schemas/expconf"
	"github.com/determined-ai/determined/proto/pkg/apiv1"
)

// nolint: exhaustruct
func createTestExpForRun(
	t *testing.T, api *apiServer, curUser model.User, projectID int, labels ...string,
) *model.Experiment {
	labelMap := make(map[string]bool)
	for _, l := range labels {
		labelMap[l] = true
	}

	activeConfig := schemas.Merge(minExpConfig, expconf.ExperimentConfig{
		RawLabels:      labelMap,
		RawDescription: ptrs.Ptr("desc"),
		RawName:        expconf.Name{RawString: ptrs.Ptr("name")},
	})
	activeConfig = schemas.WithDefaults(activeConfig)
	exp := &model.Experiment{
		JobID:     model.JobID(uuid.New().String()),
		State:     model.PausedState,
		OwnerID:   &curUser.ID,
		ProjectID: projectID,
		StartTime: time.Now(),
		Config:    activeConfig.AsLegacy(),
	}
	require.NoError(t, api.m.db.AddExperiment(exp, []byte{10, 11, 12}, activeConfig))

	// Get experiment as our API mostly will to make it easier to mock.
	exp, err := db.ExperimentByID(context.TODO(), exp.ID)
	require.NoError(t, err)
	return exp
}

func TestSearchRunsSort(t *testing.T) {
	api, curUser, ctx := setupAPITest(t, nil)
	_, projectIDInt := createProjectAndWorkspace(ctx, t, api)
	projectID := int32(projectIDInt)

	// Empty response causes no errors.
	req := &apiv1.SearchRunsRequest{
		ProjectId: &projectID,
		Sort:      ptrs.Ptr("id=asc"),
	}
	resp, err := api.SearchRuns(ctx, req)
	require.NoError(t, err)
	require.Len(t, resp.Runs, 0)

	hyperparameters := map[string]any{"global_batch_size": 1}

	exp := createTestExpForRun(t, api, curUser, projectIDInt)

	task := &model.Task{TaskType: model.TaskTypeTrial, TaskID: model.NewTaskID()}
	require.NoError(t, db.AddTask(ctx, task))
	require.NoError(t, db.AddTrial(ctx, &model.Trial{
		State:        model.PausedState,
		ExperimentID: exp.ID,
		StartTime:    time.Now(),
		HParams:      hyperparameters,
	}, task.TaskID))

	resp, err = api.SearchRuns(ctx, req)
	require.NoError(t, err)
	require.Len(t, resp.Runs, 1)

	hyperparameters2 := map[string]any{"global_batch_size": 2}

	// Add second experiment
	exp2 := createTestExpForRun(t, api, curUser, projectIDInt)

	task2 := &model.Task{TaskType: model.TaskTypeTrial, TaskID: model.NewTaskID()}
	require.NoError(t, db.AddTask(ctx, task2))
	require.NoError(t, db.AddTrial(ctx, &model.Trial{
		State:        model.PausedState,
		ExperimentID: exp2.ID,
		StartTime:    time.Now(),
		HParams:      hyperparameters2,
	}, task2.TaskID))

	// Sort by start time
	resp, err = api.SearchRuns(ctx, &apiv1.SearchRunsRequest{
		ProjectId: req.ProjectId,
		Sort:      ptrs.Ptr("startTime=asc"),
	})

	require.NoError(t, err)
	require.Equal(t, int32(exp.ID), *resp.Runs[0].ExperimentId)
	require.Equal(t, int32(exp2.ID), *resp.Runs[1].ExperimentId)

	// Sort by hyperparameter
	resp, err = api.SearchRuns(ctx, &apiv1.SearchRunsRequest{
		ProjectId: req.ProjectId,
		Sort:      ptrs.Ptr("hp.global_batch_size=desc"),
	})

	require.NoError(t, err)
	require.Equal(t, int32(exp2.ID), *resp.Runs[0].ExperimentId)
	require.Equal(t, int32(exp.ID), *resp.Runs[1].ExperimentId)
}

func TestSearchRunsFilter(t *testing.T) {
	api, curUser, ctx := setupAPITest(t, nil)
	_, projectIDInt := createProjectAndWorkspace(ctx, t, api)
	projectID := int32(projectIDInt)

	// Empty response causes no errors.
	req := &apiv1.SearchRunsRequest{
		ProjectId: &projectID,
		Sort:      ptrs.Ptr("id=asc"),
	}
	resp, err := api.SearchRuns(ctx, req)
	require.NoError(t, err)
	require.Len(t, resp.Runs, 0)

	hyperparameters := map[string]any{"global_batch_size": 1}

	exp := createTestExpForRun(t, api, curUser, projectIDInt)

	task := &model.Task{TaskType: model.TaskTypeTrial, TaskID: model.NewTaskID()}
	require.NoError(t, db.AddTask(ctx, task))
	require.NoError(t, db.AddTrial(ctx, &model.Trial{
		State:        model.PausedState,
		ExperimentID: exp.ID,
		StartTime:    time.Now(),
		HParams:      hyperparameters,
	}, task.TaskID))

	resp, err = api.SearchRuns(ctx, req)
	require.NoError(t, err)
	require.Len(t, resp.Runs, 1)

	hyperparameters2 := map[string]any{"global_batch_size": 2}

	// Add second experiment
	exp2 := createTestExpForRun(t, api, curUser, projectIDInt)

	task2 := &model.Task{TaskType: model.TaskTypeTrial, TaskID: model.NewTaskID()}
	require.NoError(t, db.AddTask(ctx, task2))
	require.NoError(t, db.AddTrial(ctx, &model.Trial{
		State:        model.PausedState,
		ExperimentID: exp2.ID,
		StartTime:    time.Now(),
		HParams:      hyperparameters2,
	}, task2.TaskID))

	// Filter by experiment id
	filter := fmt.Sprintf(`{"filterGroup":{"children":[{"columnName":"experimentId","kind":"field",`+
		`"location":"LOCATION_TYPE_RUN","operator":"=","type":"COLUMN_TYPE_NUMBER","value":%d}],`+
		`"conjunction":"and","kind":"group"},"showArchived":false}`, int32(exp2.ID))
	require.NoError(t, err)
	resp, err = api.SearchRuns(ctx, &apiv1.SearchRunsRequest{
		ProjectId: req.ProjectId,
		Filter:    ptrs.Ptr(filter),
	})

	require.NoError(t, err)
	require.Len(t, resp.Runs, 1)

	// Filter by hyperparameter
	resp, err = api.SearchRuns(ctx, &apiv1.SearchRunsRequest{
		ProjectId: req.ProjectId,
		Filter: ptrs.Ptr(`{"filterGroup":{"children":[{"columnName":"hp.global_batch_size","kind":"field",` +
			`"location":"LOCATION_TYPE_RUN_HYPERPARAMETERS","operator":"<=","type":"COLUMN_TYPE_NUMBER","value":1}],` +
			`"conjunction":"and","kind":"group"},"showArchived":false}`),
	})

	require.NoError(t, err)
	require.Len(t, resp.Runs, 1)
}
