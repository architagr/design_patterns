import { IPhone } from "./IPhone";
import { PhoneType, StorageOptions } from "./enum";

export class IPhonePlus extends IPhone {

    constructor(storage: StorageOptions) {
        super();
        this.storage = storage;
        this.screenSize = 5;
        this.cameraPixels = 12;
        this.phoneType = PhoneType.Plus;
    }
}