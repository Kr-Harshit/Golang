import React, { useEffect, useReducer } from "react";
import { Typography } from "@mui/material";

import NavMenu from "../NavMenu";
import CoffeeList from "./CoffeeList";
import { Product, Actions, ActionType } from "../../model";
import axios from "../../axios";

const productReducer = (products: Product[], action: Actions) => {
	switch (action.type) {
		case ActionType.GET:
			// setting fetch data to products
			if (action.payload)
				return Object.entries(action.payload).map(([_, product]) => {
					return {
						id: product["_id"],
						name: product.name,
						description: product.description,
						sku: product.sku,
						price: product.price,
					};
				});
			else return products;
		default:
			return products;
	}
};

const Menu: React.FC = () => {
	const [products, productsDispatch] = useReducer(productReducer, []);

	useEffect(() => {
		// fetching data when page render for first time
		axios.get("/products").then((res) => {
			productsDispatch({ type: ActionType.GET, payload: res.data });
		});
	}, []);

	return (
		<div className="Menu">
			<NavMenu />
			<Typography variant="h4" sx={{ margin: "1em 0" }}>
				Menu
			</Typography>
			<CoffeeList rows={products} />
		</div>
	);
};

export default Menu;
