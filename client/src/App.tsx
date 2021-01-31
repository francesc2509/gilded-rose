import { ApolloClient, ApolloProvider, InMemoryCache } from '@apollo/client';
import React from 'react';
import { BrowserRouter as Router} from 'react-router-dom';
import './App.css';
import { Header, Main } from './layout';

const client = new ApolloClient({
  uri: 'http://127.0.0.1:8080/graphql',
  cache: new InMemoryCache()
});

function App() {
  return (
    <ApolloProvider client={client}>
      <Router>
        <div className="App">
          <div className="App-header">
            <Header />
          </div>
          <div className="App-main">
            <Main />
          </div>
        </div>
      </Router>
    </ApolloProvider>
  );
}

export default App;
