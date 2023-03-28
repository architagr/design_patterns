import { IPhone } from "./IPhone";
import { PhoneType, SimType, StorageOptions } from "./enum";

export class IPhonePro extends IPhone {

    constructor(storage: StorageOptions, canTurnOffShutterSound: boolean, simTypes: SimType[]) {
        super();
        this.storage = storage;
        this.screenSize = 6;
        this.cameraPixels = 48;
        this.phoneType = PhoneType.Pro;
        this.canTurnOffShutterSound = canTurnOffShutterSound;
        this.shutterSound = true;
        this.sims = simTypes;
    }
}