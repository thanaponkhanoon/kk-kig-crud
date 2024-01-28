import { CcodeInterface } from "./Iccode";
import { CusdetailInterface } from "./Icusdetail";

export interface CustomerInterface {
    ID?: number;
    Name?: string;
    CcodeID?: number;
    Ccode?: CcodeInterface;
    CusdetailID?: number;
    Cusdetail?: CusdetailInterface;
}