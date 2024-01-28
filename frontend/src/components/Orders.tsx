import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import MenuItem from "@mui/material/MenuItem"
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from "@mui/material/TextField";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";

import { PcodeInterface } from "../interfaces/Ipcode";
import { DetailInterface } from "../interfaces/Idetail";
import { ProductInteface } from "../interfaces/Iproduct";

import { type } from "os";

import { GetPcode, GetDetail } from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,

    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

interface Product {
    Cancle: () => void;
    Data: ProductInteface | undefined;
}

function Orders({ Cancle, Data }: Product) {
    const [pcodes, setPcodes] = useState<PcodeInterface[]>([]);
    const [details, setDetails] = useState<DetailInterface[]>([]);
    const [product, setProduct] = useState<Partial<ProductInteface>>({
        Time: new Date(),
        ID: Data?.ID,
        PcodeID: Data?.PcodeID,
        DetailID: Data?.DetailID,
        Amount: Data?.Amount,
        Price: Data?.Price,
    })
    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);
    const [errorMessage, setErrorMessage] = useState("");
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
    const handleChange = (event: SelectChangeEvent<number>) => {
        const name = event.target.name as keyof typeof product;
        setProduct({
            ...product,
            [name]: event.target.value,
        });
    };

    const handleInputChange = (
        event: React.ChangeEvent<{ id?: string; value: any }>
    ) => {
        const id = event.target.id as keyof typeof product;
        const { value } = event.target;
        setProduct({ ...product, [id]: value });
    };

    const getPcode = async () => {
        let res = await GetPcode();
        if (res) {
            setPcodes(res);
        }
    };

    const getDetail = async () => {
        let res = await GetDetail();
        if (res) {
            setDetails(res);
        }
    };

    useEffect(() => {
        getPcode();
        getDetail();
    }, []);

    const convertType = (data: string | number | undefined) => {
        let val = typeof data === "string" ? parseInt(data) : data;
        return val;
    };

    async function submit() {
        let data = {
            PcodeID: product.PcodeID,
            DetailID: product.DetailID,
            Name: product.Name,
            Amount: convertType(product.Amount),
            Price: convertType(product.Price),
            Time: product.Time,
        };
        console.log(data);

        const apiUrl = "http://localhost:8080/products";
        const requestOptions = {
            method: "PATCH",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },

            body: JSON.stringify(data),
        };

        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                console.log(res);
                if (res.data) {
                    console.log("บันทึกได้");
                    setSuccess(true);
                    window.location.reload();
                    setErrorMessage("");
                } else {
                    console.log("บันทึกไม่ได้");
                    setError(true);
                    setErrorMessage(res.error);
                }
            });
    }

    return (
        <div>
            <Container maxWidth="md">
                <Snackbar
                    id="success"
                    open={success}
                    autoHideDuration={6000}
                    onClose={handleClose}
                    anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                >
                    <Alert onClose={handleClose} severity="success">
                        บันทึกสำเร็จ
                    </Alert>
                </Snackbar>
                <Snackbar
                    id="error"
                    open={error}
                    autoHideDuration={6000}
                    onClose={handleClose}
                >
                    <Alert onClose={handleClose} severity="error">
                        บันทึกไม่สำเร็จ: {errorMessage}
                    </Alert>
                </Snackbar>

                <Paper>
                    <Box
                        display="flex"
                        sx={{
                            marginTop: 2,
                        }}
                    >
                        <Box sx={{ paddingX: 2, paddingY: 1 }}>
                            <Typography
                                component="h1"
                                variant="h6"
                                color="primary"
                                gutterBottom
                            >
                                แก้ไขรายการสั่งซื้อสินค้า
                            </Typography>
                        </Box>
                    </Box>

                    <Divider />
                    <Grid container spacing={3} sx={{ padding: 2 }}>
                        <Grid item xs={6}>
                            <FormControl fullWidth variant="outlined">
                                <b>
                                    <p>รหัสสินค้า</p>
                                </b>

                                <Select
                                    value={product.PcodeID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "PcodeID", //เอาไว้เข้าถึงข้อมูลแพลนนิ่งไอดี
                                    }}
                                >
                                    <option aria-label="None" value="">
                                        กรุณาเลือกรหัสสินค้า
                                    </option>
                                    {pcodes.map(
                                        (
                                            item: PcodeInterface //map
                                        ) => (
                                            <MenuItem value={item.ID} key={item.ID}>
                                                {item.Name}
                                            </MenuItem> //key ไว้อ้างอิงว่าที่1ชื่อนี้ๆๆ value: เก็บค่า
                                        )
                                    )}
                                </Select>
                            </FormControl>
                        </Grid>

                        <Grid item xs={6}>
                            <FormControl fullWidth variant="standard">
                                <p>รายการสั่งซื้อสินค้า</p>

                                <Select
                                    value={product.DetailID}
                                    onChange={handleChange}
                                    inputProps={{
                                        name: "DetailID", //เอาไว้เข้าถึงข้อมูลแพลนนิ่งไอดี
                                    }}
                                >
                                    <option aria-label="None" value="">
                                        กรุณาเลือกรายลัเอียดสินค้า
                                    </option>
                                    {details.map(
                                        (
                                            item: DetailInterface //map
                                        ) => (
                                            <MenuItem value={item.ID} key={item.ID}>
                                                {item.Name}
                                            </MenuItem> //key ไว้อ้างอิงว่าที่1ชื่อนี้ๆๆ value: เก็บค่า
                                        )
                                    )}
                                </Select>
                            </FormControl>
                        </Grid>

                        <Grid item xs={6}>
                            <FormControl fullWidth variant="outlined">
                                <p>วันที่และเวลา</p>
                                <LocalizationProvider dateAdapter={AdapterDateFns}>
                                    <DatePicker
                                        value={product.Time}
                                        onChange={(newValue) => {
                                            setProduct({
                                                ...product,
                                                Time: newValue,
                                            });
                                        }}
                                        renderInput={(params) => <TextField {...params} />}
                                    />
                                </LocalizationProvider>
                            </FormControl>
                        </Grid>
                        <Grid item xs={6}>
                            <FormControl fullWidth variant="outlined">
                                <p>ชื่อลูกค้า</p>
                                <TextField
                                    id="Name"
                                    value={product.Name ?? ""}
                                    onChange={handleInputChange}
                                />
                            </FormControl>
                        </Grid>
                        <Grid item xs={6}>
                            <FormControl fullWidth variant="outlined">
                                <p>จำนวน</p>
                                <TextField
                                    id="Amount"
                                    type="number"
                                    value={product.Amount ?? 0}
                                    onChange={handleInputChange}
                                />
                            </FormControl>
                        </Grid>
                        <Grid item xs={6}>
                            <FormControl fullWidth variant="outlined">
                                <p>ราคา</p>
                                <TextField
                                    id="Price"
                                    type="number"
                                    value={product.Price ?? 0}
                                    onChange={handleInputChange}
                                />
                            </FormControl>
                        </Grid>

                        <Grid item xs={12}>
                            <Button variant="contained" onClick={Cancle}>
                                กลับ
                            </Button>

                            <Button
                                style={{ float: "right" }}
                                onClick={submit}
                                variant="contained"
                                color="success"
                            >
                                บันทึกการแก้ไข
                            </Button>
                        </Grid>
                    </Grid>
                </Paper>
            </Container>
        </div>
    );
}

export default Orders;