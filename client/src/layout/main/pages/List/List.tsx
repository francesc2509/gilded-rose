import { gql, useMutation, useQuery } from '@apollo/client';
import React from 'react';
import { Item } from '../../../../entities';
import { ListItem } from './components';

import './List.css'

const GET_ITEM_QUERY = gql`
    query get_item {
        get_item {
            id,
            name,
            sellIn,
            quality
        }
    }
`;

const UPDATE_ITEM_INVENTORY_QUERY = gql`
    mutation update_item_inventory($days: Int!) {
        update_item_inventory(days: $days) {
            id,
            name,
            sellIn,
            quality
        }
    }
`;

const List: React.FunctionComponent = () => {
    const {loading, data, error} = useQuery(GET_ITEM_QUERY);
    const [ updateInventory ] = useMutation(UPDATE_ITEM_INVENTORY_QUERY);

    if (loading) {
        return (
            <h2>Loading</h2>
        );
    }

    if (error) {
        return (<h2>{error.message}</h2>);
    }

    const items = data.get_item;

    return (
        <div>
        <table>
            <thead>
                <tr>
                    <th>Name</th>
                    <th>Sell in (days)</th>
                    <th>Quality</th>
                </tr>
            </thead>
            <tbody>
                {
                    items && items.filter((item: Item) => item).map((item: Item) => {
                        return (<ListItem item={item} key={item.id} />);
                    })
                }
            </tbody>
        </table>
        <button onClick={(e) => {
                updateInventory({variables: { days: 1 }});
            }}>
            Update
        </button>
        </div>
    );
}

export default List;