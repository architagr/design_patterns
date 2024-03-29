import { IPhone } from "./IPhone";
import { PhoneType, SimType, StorageOptions } from "./enum";

export class IPhonePlus extends IPhone {

    constructor(storage: StorageOptions, canTurnOffShutterSound: boolean, simTypes: SimType[]) {
        super();
        this.storage = storage;
        this.screenSize = 5;
        this.cameraPixels = 12;
        this.phoneType = PhoneType.Plus;
        this.canTurnOffShutterSound = canTurnOffShutterSound;
        this.shutterSound = true;
        this.sims = simTypes;
    }
}