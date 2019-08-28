import React, { Component } from 'react';
import logo from './logo.svg';
import { GoogleLogin } from 'react-google-login';
import './App.css';

class App extends Component {
  constructor(props){
    super(props)
    this.state = {
      authorized: false,
      failed: false
    }
    this.authorize = this.authorize.bind(this)
    this.authorizeFailed = this.authorizeFailed.bind(this)
  }
  render(){return (
    <div className="App">
      <header className="App-header">
      { this.state.authorized ? "" :
      <GoogleLogin
        clientId="863546706494-hv1o3g1v7ao8londkl58b8ogpifvjvti.apps.googleusercontent.com"
        buttonText="Login"
        onSuccess={this.authorize}
        onFailure={this.authorizeFailed}
        cookiePolicy={'single_host_origin'}
      />}
      { !this.state.failed && this.state.authorized ? "The answer to all of your questions is.... 42!" : ""}
      { this.state.failed ? "Not authorized :(" : ""}
      </header>
    </div>
  );}

  authorize(response){
    console.log("AUTHORIZED!!");
    this.setState({
      authorized: true,
      failed: false
    })
  }

  authorizeFailed(response){
    console.log("FAILED!!")
    this.setState({
      authorized: false,
      failed: true
    })
  }
}

export default App;
