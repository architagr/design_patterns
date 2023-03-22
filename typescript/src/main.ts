import { IPhone } from "./creational/factory_pattern/IPhone";
import { IPhonePlus } from "./creational/factory_pattern/IPhonePlus";
import { IPhonePro } from "./creational/factory_pattern/IPhonePro";
import { PhoneType, StorageOptions } from "./creational/factory_pattern/enum";
import { Store } from "./creational/factory_pattern/store";

const main = (): void => {
    let pro128: IPhone = new IPhonePro(StorageOptions.Gb128);
    console.log("*** Pro ***");

    console.log(`Screen size ${pro128.getScreenSize()}`);
    console.log(`Screen size ${pro128.getCameraPixels()}`);

    let plus128: IPhone = new IPhonePlus(StorageOptions.Gb128);
    console.log("*** Plus ***");
    console.log(`Screen size ${plus128.getScreenSize()}`);
    console.log(`Screen size ${plus128.getCameraPixels()}`);

    console.log("*** Pro from store ***");

    let store: Store = new Store();
    let phone: IPhone | null = store.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
    console.log(`Screen size ${phone!.getScreenSize()}`);
    console.log(`Screen size ${phone!.getCameraPixels()}`);
}

main();