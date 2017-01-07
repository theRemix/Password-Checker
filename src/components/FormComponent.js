import Inferno from 'inferno';

export function FormComponent() {
	return (
		<fieldset>
      <div className="field">
        <label htmlFor="username">Enter Your Username</label>
        <input id="username" type="text" name="username" />
      </div>
      <div className="field">
        <label htmlFor="password">Enter Your Password</label>
        <input id="password" type="text" name="password" />
      </div>
      <div className="field">
        <button type="button">Check</button>
      </div>
		</fieldset>
	);
}
