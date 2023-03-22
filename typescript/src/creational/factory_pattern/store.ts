import { IPhone } from "./IPhone";
import { IPhonePlus } from "./IPhonePlus";
import { IPhonePro } from "./IPhonePro";
import { PhoneType, StorageOptions } from "./enum";

export class Store {

    public CreateOrder(type: PhoneType, storage: StorageOptions): IPhone | null {
        let phone: IPhone | null = null;
        if (type == PhoneType.Pro) {
            phone = new IPhonePro(storage);
        } else if (type == PhoneType.Plus) {
            phone = new IPhonePlus(storage);
        } else {
            throw new Error("Invalid Phone type")
        }

        phone!.InstallSoftwares();
        phone!.Package();
        return phone;
    }
}