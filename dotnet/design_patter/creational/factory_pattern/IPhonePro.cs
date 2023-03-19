namespace creational.factory_pattern{
    public class IPhonePro: IPhone{

        public IPhonePro(StorageOptions storage ){
            this.storage = storage;
            this.screenSize = 6;
            this.cameraPixels = 48;
            this.phoneType = PhoneType.Pro;
        }
        
    }
}