import Inferno from 'inferno';
import Component from 'inferno-component';

export class FormComponent extends Component{

 constructor() {
    super();
    this.state = {
      username: null,
      password: null,
      errors: null
    };
    this.updateUsername = this.updateUsername.bind(this);
    this.updatePassword = this.updatePassword.bind(this);
    this.submit = this.submit.bind(this);
  }

  updateUsername(event) {
    this.setState({username: event.target.value});
  }

  updatePassword(event) {
    this.setState({password: event.target.value});
  }

  submit(event) {
    event.preventDefault();
    console.log(event);
    // handle form submission
    const req = new Request("/api/check", { method: "POST", body: `{"username": ${this.state.username}, "password": ${this.state.password}}` });
    fetch(req)
      .then(res => res.json())
      .then(res => {
        let result = res.success ?
          "Congratulations, this is your correct password!" :
          "This is not your password!";
        this.setState({ result });
      })
      .catch(err => {
        let errors = err instanceof SyntaxError ?
          "Server found a broken" :
          err.toString();
        this.setState({ errors });
      });
  }

  render() {
    return (
      <div>
        <h2>{this.state.result}</h2>
        <form action="#" onSubmit={this.submit}>
          <div className="field">
            <label htmlFor="username">Enter Your Username</label>
            <input id="username" type="text" name="username" onChange={this.updateUsername} />
          </div>
          <div className="field">
            <label htmlFor="password">Enter Your Password</label>
            <input id="password" type="password" name="password" onChange={this.updatePassword} />
          </div>
          <div className="field">
            <button type="submit">Check</button>
          </div>
        </form>
        <div className="error">{this.state.errors}</div>
      </div>
    );
  }
}
