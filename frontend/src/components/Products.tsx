import React, { useState, useEffect, useCallback } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { ProductInteface } from "../interfaces/Iproduct";
import { GetProduct } from "../services/HttpClientService";
import moment from "moment";
import EditIcon from "@mui/icons-material/Edit";
import DeleteForeverIcon from "@mui/icons-material/DeleteForever";
import Snackbar from "@mui/material/Snackbar";
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogTitle from "@mui/material/DialogTitle";
import Orders from "./Orders";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,

    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Products() {
    const [products, setProducts] = useState<ProductInteface[]>([]);
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [opendelete, setOpenDelete] = useState(false);
    const [openedit, setOpenEdit] = useState(false);
    const [selectcellData, setSelectcellData] = useState<ProductInteface>();

    const [selectcell, setSelectCell] = useState(Number);
    const handleCellFocus = useCallback(
        (event: React.FocusEvent<HTMLDivElement>) => {
            const row = event.currentTarget.parentElement;
            const id = row!.dataset.id!;
            const b = products.filter((v) => Number(v.ID) == Number(id));
            const field = event.currentTarget.dataset.field!;
            console.log(b[0]);
            setSelectcellData(b[0]);
            setSelectCell(Number(id));
        },
        [products]
    );
    const handleClose = (
        event?: React.SyntheticEvent | Event,

        reason?: string
    ) => {
        console.log(reason);
        if (reason === "clickaway") {
            return;
        }

        setSuccess(false);

        setError(false);
    };
    const handleClickDelete = () => {
        DeleteProduct(selectcell);
        setOpenDelete(false);
    };
    const handleDelete = () => {
        setOpenDelete(true);
    };
    const handleEdit = () => {
        setOpenEdit(true);
    };
    const handleDeleteClose = () => {
        setOpenDelete(false);
    };
    const handleEditClose = () => {
        setOpenEdit(false);
    };

    useEffect(() => {
        getProducts();
    }, []);

    const getProducts = async () => {
        let res = await GetProduct();
        console.log(res);
        if (res) {
            setProducts(res);
        }
    };

    const DeleteProduct = async (id: Number) => {
        const apiUrl = `http://localhost:8080/products/${id}`;
        const requestOptions = {
            method: "DELETE",

            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`, //การยืนยันตัวตน
                "Content-Type": "application/json",
            },
        };

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())

            .then((res) => {
                console.log(res);

                if (res) {
                    setSuccess(true);
                    window.location.reload();
                } else {
                    setError(true);
                }
            });
    };

    const columns: GridColDef[] = [
        {
            field: "Pcode",
            headerName: "รหัสสินค้า",
            width: 100,
            valueGetter: (params) => params.row.Pcode.Name,
        },
        {
            field: "Name",
            headerName: "ชื่อลูกค้า",
            width: 100,
            valueFormatter: (params) => params.value.Name,
        },
        { field: "ID", headerName: "ลำดับ", width: 50 },
        {
            field: "Detail",
            headerName: "รายละเอียดสินค้า",
            width: 130,
            valueGetter: (params) => params.row.Detail.Name,
        },
        {
            field: "Price",
            headerName: "ราคา",
            width: 70,
            valueFormatter: (params) => params.value.Price,
        },
        {
            field: "Amount",
            headerName: "จำนวน",
            width: 70,
            valueFormatter: (params) => params.value.Amount,
        },
        {
            field: "Time", headerName: "วันที่และเวลา", width: 150,
            valueFormatter: (params) => moment(params.value.Time)
        },
        {
            field: "Action", headerName: "จัดการ", width: 200,
            renderCell: () => (
                <div>
                    &nbsp;
                    <Button
                        onClick={handleEdit}
                        variant="contained"
                        size="small"
                        startIcon={<EditIcon />}
                        color="success"
                    >
                        แก้ไข
                    </Button>
                    &nbsp;&nbsp;&nbsp;
                    <Button
                        onClick={handleDelete}
                        variant="contained"
                        size="small"
                        startIcon={<DeleteForeverIcon />}
                        color="error"
                    >
                        ลบ
                    </Button>
                    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
                </div>
            ),
        },
    ]

    return (
        <div>
            <Container maxWidth="md">
                <Snackbar
                    open={success}
                    autoHideDuration={6000}
                    onClose={handleClose}
                    anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                >
                    <Alert onClose={handleClose} severity="success">
                        ลบข้อมูลสำเร็จ
                    </Alert>
                </Snackbar>

                <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                    <Alert onClose={handleClose} severity="error">
                        ลบข้อมูลไม่สำเร็จ
                    </Alert>
                </Snackbar>
                <Dialog
                    open={opendelete}
                    onClose={handleDeleteClose}
                    aria-labelledby="alert-dialog-title"
                    aria-describedby="alert-dialog-description"
                >
                    <DialogTitle id="alert-dialog-title">
                        {"คุณต้องการลบรายการสั่งซื้อใช่หรือไม่?"}
                    </DialogTitle>

                    <DialogActions>
                        <Button onClick={handleDeleteClose}>ยกเลิก</Button>
                        <Button onClick={handleClickDelete} autoFocus>
                            ตกลง
                        </Button>
                    </DialogActions>
                </Dialog>

                <Dialog
                    open={openedit}
                    onClose={handleEditClose}
                    aria-labelledby="alert-dialog-title"
                    aria-describedby="alert-dialog-description"
                >
                    <DialogActions>
                        <Orders
                            Cancle={handleEditClose}
                            Data={selectcellData}
                        />
                    </DialogActions>
                </Dialog>

                <Box
                    display="flex"
                    sx={{
                        marginTop: 2,
                    }}
                >
                    <Box flexGrow={1}>
                        <Typography
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            รายการสั่งซื้อสินค้า
                        </Typography>
                    </Box>
                </Box>
                <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
                    <DataGrid
                        rows={products}
                        getRowId={(row) => row.ID}
                        columns={columns}
                        pageSize={10}
                        componentsProps={{
                            cell: {
                                onFocus: handleCellFocus,
                            },
                        }}
                        rowsPerPageOptions={[10]}
                    />
                </div>
                <br></br>
                <div>
                    <Box>
                        <Button
                            style={{ float: "right" }}
                            component={RouterLink}
                            to="/customers"
                            variant="contained"
                            color="primary"
                        >
                            เพิ่มรายการสั่งซื้อสินค้า
                        </Button>
                    </Box>
                </div>
            </Container>
        </div>
    );
}

export default Products;