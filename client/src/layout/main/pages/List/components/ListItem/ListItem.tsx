import React, { FunctionComponent } from 'react';
import { Item } from '../../../../../../entities';

interface ListItemProps {
    item: Item;
}

const ListItem: FunctionComponent<ListItemProps> = ({ item }) => {
    return (
        <tr>
            <td>{item.name}</td>
            <td>{item.sellIn}</td>
            <td>{item.quality}</td>
        </tr>
    );
};

export default ListItem;