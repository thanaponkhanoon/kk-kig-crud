import React, { useEffect, useState } from "react";
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

import { GetPcode, GetDetail, CreateProduct } from "../services/HttpClientService";

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
    props,
    ref
) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
});

function Customers() {
    const [pcodes, setPcodes] = useState<PcodeInterface[]>([]);
    const [details, setDetails] = useState<DetailInterface[]>([]);
    const [product, setProduct] = useState<Partial<ProductInteface>>({
        Time: new Date(),
    })

    const [success, setSuccess] = useState(false);
    const [error, setError] = useState(false);

    const handleClose = (
        event?: React.SyntheticEvent | Event,
        reason?: string
    ) => {
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

        let res = await CreateProduct(data as ProductInteface);
        if (res) {
            setSuccess(true);
        } else {
            setError(true);
        }
    }

    return (
        <Container maxWidth="md">
            <Snackbar
                open={success}
                autoHideDuration={3000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="success">
                    บันทึกรายการสั่งซื้อาสำเร็จ
                </Alert>
            </Snackbar>
            <Snackbar
                open={error}
                autoHideDuration={6000}
                onClose={handleClose}
                anchorOrigin={{ vertical: "top", horizontal: "center" }}
            >
                <Alert onClose={handleClose} severity="error">
                    บันทึกรายการสั่งซื้อไม่สำเร็จ
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
                            component="h2"
                            variant="h6"
                            color="primary"
                            gutterBottom
                        >
                            บันทึกรายการสั่งซื้อ
                        </Typography>
                    </Box>
                </Box>
                <Divider />
                <Grid container spacing={3} sx={{ padding: 2 }}>
                    <Grid item xs={6}>
                        <FormControl fullWidth variant="outlined">
                            <p>รหัสสินค้า</p>
                            <Select
                                value={product.PcodeID ?? 0}
                                onChange={handleChange}
                                inputProps={{
                                    name: "PcodeID",
                                }}
                            >
                                <MenuItem aria-label="None" value={0}>
                                    กรุณาเลือกรหัสสินค้า
                                </MenuItem>
                                {pcodes.map((item: PcodeInterface) => (
                                    <MenuItem value={item.ID} key={item.ID}>
                                        {item.Name}
                                    </MenuItem>
                                ))}
                            </Select>
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
                            <p>รายละเอียดสินค้า</p>
                            <Select
                                value={product.DetailID ?? 0}
                                onChange={handleChange}
                                inputProps={{
                                    name: "DetailID",
                                }}
                            >
                                <MenuItem aria-label="None" value={0}>
                                    กรุณาเลือกรายลัเอียดสินค้า
                                </MenuItem>
                                {details.map((item: DetailInterface) => (
                                    <MenuItem value={item.ID} key={item.ID}>
                                        {item.Name}
                                    </MenuItem>
                                ))}
                            </Select>
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
                    <Grid item xs={12}>
                        <Button
                            component={RouterLink}
                            to="/products"
                            variant="contained"
                            color="inherit"
                        >
                            กลับ
                        </Button>
                        <Button
                            style={{ float: "right" }}
                            onClick={submit}
                            variant="contained"
                            color="primary"
                        >
                            บันทึก
                        </Button>
                    </Grid>
                </Grid>
            </Paper>
        </Container>
    );
}

export default Customers;