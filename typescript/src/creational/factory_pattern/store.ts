import { IPhone } from "./IPhone";
import { IPhonePlus } from "./IPhonePlus";
import { IPhonePro } from "./IPhonePro";
import { PhoneType, SimType, StorageOptions } from "./enum";


export abstract class Store {

    public CreateOrder(type: PhoneType, storage: StorageOptions): IPhone | null {
        let phone = this.createPhone(type, storage);
        if (phone) {
            phone.InstallSoftwares();
            phone.Package();
        }
        return phone;
    }
    public abstract createPhone(type: PhoneType, storage: StorageOptions): IPhone | null;
}

export class IndiaStore extends Store {
    private simTypes: SimType[] = [SimType.Esim, SimType.PhysicalSim];
    private canTurnOffShutterSound: boolean = true;


    public override createPhone(type: PhoneType, storage: StorageOptions): IPhone | null {
        let phone: IPhone | null = null;
        if (type == PhoneType.Pro) {
            phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
        } else if (type == PhoneType.Plus) {
            phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
        }
        return phone;
    }
}
export class JapanStore extends Store {
    private simTypes: SimType[] = [SimType.PhysicalSim, SimType.PhysicalSim];
    private canTurnOffShutterSound: boolean = false;

    public override createPhone(type: PhoneType, storage: StorageOptions): IPhone | null {
        let phone: IPhone | null = null;
        if (type == PhoneType.Pro) {
            phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
        } else if (type == PhoneType.Plus) {
            phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
        }
        return phone;
    }
}

export class USAStore extends Store {
    private simTypes: SimType[] = [SimType.Esim, SimType.Esim];
    private canTurnOffShutterSound: boolean = true;
    public override createPhone(type: PhoneType, storage: StorageOptions): IPhone | null {
        let phone: IPhone | null = null;
        if (type == PhoneType.Pro) {
            phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
        } else if (type == PhoneType.Plus) {
            phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
        }
        return phone;
    }
}