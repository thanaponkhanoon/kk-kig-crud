import { PcodeInterface } from "./Ipcode";
import { DetailInterface } from "./Idetail";

export interface ProductInteface {
    ID?: number;
    Name?: string;
    Price?: number;
    Amount?: number;
    PcodeID?: number;
    Pcode?: PcodeInterface;
    DetailID?: number;
    Detail?: DetailInterface;
    Time: Date | null;
}