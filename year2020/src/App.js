import React from 'react';
import logo from './logo.svg';
import Countdown from 'react-countdown-now';

import './App.css';

function App() {
  
    // TImer completion
    const Completionist = () => <span>WOW, IT'S 2020 ALREADY!</span>;
  
    // Renderer callback with condition
    const renderer = ({ days, hours, minutes, seconds, completed }) => {
      if (completed) {
        // Render a completed state
        return <Completionist />;
      } else {
        // Render a countdown
        return <span>{days*24*60*60 + hours*60*60 + minutes * 60 + seconds}</span>;
      }
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>SECONDS LEFT UNTIL YEAR 2020:</p>
        <p><Countdown date={new Date(2020, 1, 1)} renderer={renderer}/></p>
        <p>(SPEND THAT TIME REASONABLY, BRO)</p>
      </header>
    </div>
  );
}

export default App;
