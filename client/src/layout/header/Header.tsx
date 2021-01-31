import React from 'react';
import { Link, Route } from 'react-router-dom';

import './Header.css'
import logo from './rose.png'

const Header: React.FunctionComponent = () => {
    return (
        <header>
            <h1>
                <Link to="/">
                    <img src={logo} alt="Home" />
                </Link>
                Gilded Rose Inn
            </h1>
        </header>
    );
}

export default Header;