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

type ProvidedMenuProp = {|title: string, icon: string|};
const menuSkeletons: Array<{|id: string, menu: ProvidedMenuProp|}> = [
	{
		id:   'home',
		menu: {
			title: 'Home',
			icon:  'home',
		},
	}, {
		id:   'chain',
		menu: {
			title: 'Chain',
			icon:  'link',
		},
	}, {
		id:   'txpool',
		menu: {
			title: 'TxPool',
			icon:  'credit-card',
		},
	}, {
		id:   'network',
		menu: {
			title: 'Network',
			icon:  'globe',
		},
	}, {
		id:   'system',
		menu: {
			title: 'System',
			icon:  'tachometer',
		},
	}, {
		id:   'logs',
		menu: {
			title: 'Logs',
			icon:  'list',
		},
	},
];
export type MenuProp = {|...ProvidedMenuProp, id: string|};
// The sidebar menu and the main content are rendered based on these elements.
// Using the id is circumstantial in some cases, so it is better to insert it also as a value.
// This way the mistyping is prevented.
export const MENU: Map<string, {...MenuProp}> = new Map(menuSkeletons.map(({id, menu}) => ([id, {id, ...menu}])));

export const DURATION = 200;

export const styles = {
	light: {
		color: 'rgba(255, 255, 255, 0.54)',
	},
}
