import { PhoneType, StorageOptions } from "./enum";

export abstract class IPhone {
    protected screenSize: number = 0;
    protected cameraPixels: number = 0;
    protected phoneType: PhoneType = PhoneType.Pro;
    protected storage: StorageOptions = StorageOptions.Gb128;

    public InstallSoftwares(): boolean {
        console.log(`please wait we are updating the softwares for your ${this.phoneType}, ${this.storage}.`);
        console.log(`All softwares updated for ${this.phoneType}, ${this.storage}.`);
        return true;
    }

    public Package(): boolean {
        console.log(`${this.phoneType}, ${this.storage} is getting packed !!!`);
        console.log(`${this.phoneType}, ${this.storage} is packed !!!, PARTY`);
        return true;
    }
    public getPhoneType(): PhoneType {
        return this.phoneType;
    }
    public getStorage(): StorageOptions {
        return this.storage;
    }
    public getScreenSize(): number {
        return this.screenSize;
    }
    public getCameraPixels(): number {
        return this.cameraPixels;
    }
}