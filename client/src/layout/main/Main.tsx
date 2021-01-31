import React from 'react';
import { List } from './pages';

import './Main.css'
import { Redirect, Route, Switch } from 'react-router';

const Main: React.FunctionComponent = () => {
    return (
        <main>
            <Switch>
                <Route path="/home">
                    <List />
                </Route>
                <Redirect from="/" to="/home" />
            </Switch>
        </main>
    );
};

export default Main;