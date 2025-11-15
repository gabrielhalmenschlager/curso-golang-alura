import React, { Component } from 'react';
import axios from 'axios';
import './Personalidades.css';

export default class Personalidades extends Component {
  state = {
    personalidades: []
  };

  componentDidMount() {
    axios.get('http://localhost:8000/api/personalidades')
      .then(res => {
        this.setState({ personalidades: res.data });
      });
  }

  render() {
    return (
      <div>
        {this.state.personalidades.map((p, id) =>
          <div className="CardPersonalidades" key={id}>
            <h3>{p.nome}</h3>
            <p>{p.historia}</p>
          </div>
        )}
      </div>
    );
  }
}
