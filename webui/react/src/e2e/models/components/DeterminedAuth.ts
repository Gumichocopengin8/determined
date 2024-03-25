import { BaseComponent, NamedComponent, NamedComponentArgs } from 'e2e/models/BaseComponent';
import { ErrorComponent } from 'e2e/models/utils/error';

/**
 * Returns a representation of the DeterminedAuth component.
 *
 * @remarks
 * This constructor represents the contents in src/components/DeterminedAuth.tsx.
 *
 * @param {Object} obj
 * @param {implementsGetLocator} obj.parent - The parent used to locate this DeterminedAuth
 * @param {string} [obj.selector] - Used instead of `defaultSelector`
 */
export class DeterminedAuth extends NamedComponent {
  static override defaultSelector="div[data-test='detAuth']";
  constructor({ selector, parent }: NamedComponentArgs) {
    super({ parent: parent, selector: selector || DeterminedAuth.defaultSelector });
  }
  readonly form: BaseComponent = new BaseComponent({ parent: this, selector: "form" });
  readonly username: BaseComponent = new BaseComponent({ parent: this.form, selector: "input[data-testid='username']" });
  readonly password: BaseComponent = new BaseComponent({ parent: this.form, selector: "input[data-testid='password']" });
  readonly submit: BaseComponent = new BaseComponent({ parent: this.form, selector: "button[data-testid='submit']" });
  readonly docs: BaseComponent = new BaseComponent({ parent: this, selector: "link[data-testid='docs']" });
  // TODO consdier a BaseComponents plural class
  readonly errors: ErrorComponent = new ErrorComponent({ parent: this.root, selector: ErrorComponent.defaultSelector+ErrorComponent.selectorTopRight });
}