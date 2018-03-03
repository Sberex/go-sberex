// @flow

// This file is part of the go-sberex library. The go-sberex library is 
// free software: you can redistribute it and/or modify it under the terms 
// of the GNU Lesser General Public License as published by the Free 
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// The go-sberex library is distributed in the hope that it will be useful, 
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU Lesser 
// General Public License <http://www.gnu.org/licenses/> for more details.

import React, {Component} from 'react';

import SideBar from './SideBar';
import Main from './Main';
import type {Content} from '../types/content';

// styles contains the constant styles of the component.
const styles = {
	body: {
		display: 'flex',
		width:   '100%',
		height:  '100%',
	},
};

export type Props = {
	opened: boolean,
	changeContent: string => void,
	active: string,
	content: Content,
	shouldUpdate: Object,
};

// Body renders the body of the dashboard.
class Body extends Component<Props> {
	render() {
		return (
			<div style={styles.body}>
				<SideBar
					opened={this.props.opened}
					changeContent={this.props.changeContent}
				/>
				<Main
					active={this.props.active}
					content={this.props.content}
					shouldUpdate={this.props.shouldUpdate}
				/>
			</div>
		);
	}
}

export default Body;
