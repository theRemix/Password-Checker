import Inferno from 'inferno';
import Component from 'inferno-component';

export class FormComponent extends Component{

 constructor() {
    super();
    this.state = {
      username: null,
      password: null,
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
    // handle form submission
  }

  render() {
    return (
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
    );
  }
}
