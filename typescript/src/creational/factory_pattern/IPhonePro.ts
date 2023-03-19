import { IPhone } from "./IPhone";
import { PhoneType, StorageOptions } from "./enum";

export class IPhonePro extends IPhone {

    constructor(storage: StorageOptions) {
        super();
        this.storage = storage;
        this.screenSize = 6;
        this.cameraPixels = 48;
        this.phoneType = PhoneType.Pro;
    }
}