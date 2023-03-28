import { PhoneType, StorageOptions } from "./creational/factory_pattern/enum";
import { IndiaStore, JapanStore } from "./creational/factory_pattern/store";

const main = (): void => {
    console.log("**** India Store ****");
    let indiaStore = new IndiaStore();
    console.log("created order for pro 128 GB");

    let phone = indiaStore.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
    if (phone) {
        console.log(`Screen Size - ${phone.getScreenSize()}`);
        console.log(`Camera - ${phone.getCameraPixels()} Pixels`);
        phone.takePic();
        phone.toggleCameraShutterSound();
        phone.takePic();
        let sims = phone.getSimTypes();
        console.log(`sim types ${sims[0]}, ${sims[1]}`);

    }
    console.log("**** Japan Store ****");
    let japanStore = new JapanStore();
    console.log("created order for pro 128 GB");

    phone = japanStore.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
    if (phone) {
        console.log(`Screen Size - ${phone.getScreenSize()}`);
        console.log(`Camera - ${phone.getCameraPixels()} Pixels`);
        phone.takePic();
        phone.toggleCameraShutterSound();
        phone.takePic();

        let sims = phone.getSimTypes();
        console.log(`sim types ${sims[0]}, ${sims[1]}`);
    }
}

main();