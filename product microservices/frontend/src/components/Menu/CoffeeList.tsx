import React from "react";
import Container from "@mui/material/Container";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";

import { Product } from "../../model";

interface Props {
	rows: Product[];
}

const CoffeeList: React.FC<Props> = ({ rows }) => {
	return (
		<Container maxWidth="xl">
			<TableContainer component={Paper}>
				<Table sx={{ minWidth: 650 }} aria-label="simple table">
					<TableHead sx={{ backgroundColor: "rgba(165, 161, 161, 0.577);" }}>
						<TableRow>
							<TableCell size="small">Name</TableCell>
							<TableCell align="right">Description</TableCell>
							<TableCell align="right">SKU</TableCell>
							<TableCell align="right">Price&nbsp;($)</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{rows.map((row, idx) => (
							<TableRow
								key={row.name}
								sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
							>
								<TableCell component="th" scope="row" size="small">
									{row.name}
								</TableCell>
								<TableCell align="right">{row.description}</TableCell>
								<TableCell align="right">{row.sku}</TableCell>
								<TableCell align="right">{row.price}</TableCell>
							</TableRow>
						))}
					</TableBody>
				</Table>
			</TableContainer>
		</Container>
	);
};

export default CoffeeList;
