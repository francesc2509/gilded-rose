import { gql, useQuery } from '@apollo/client';
import React from 'react';
import { Item } from '../../../../entities';
import { ListItem } from './components';

import './List.css'

const GET_ITEMS_QUERY = gql`
    query get_item {
        get_item {
            name,
            sellIn,
            quality
        }
    }
`;

const List: React.FunctionComponent = () => {
    const {loading, data, error} = useQuery(GET_ITEMS_QUERY);

    if (loading) {
        return (
            <h2>Loading</h2>
        );
    }

    if (error) {
        return (<h2>{error.message}</h2>);
    }

    const items = data.get_item;

    console.log(items);

    return (
        <table>
            <thead>
                <th>Name</th>
                <th>Sell in (days)</th>
                <th>Quality</th>
            </thead>
            <tbody>
                {
                    items && items.filter((item: Item) => item).map((item: Item) => {
                        return (<ListItem item={item} key={item.name} />);
                    })
                }
            </tbody>
        </table>
    );
}

export default List;