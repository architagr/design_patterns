namespace creational.factory_pattern{
    public class IPhonePro: IPhone{

        public IPhonePro(StorageOptions storage ){
            this.storage = storage;
            this.screenSize = 6;
            this.cameraPixels = 48;
            this.phoneType = PhoneType.Pro;
        }
        
        public override PhoneType getPhoneType(){
            return this.phoneType;
        }
        public override StorageOptions getStorage(){
            return this.storage;
        }
        public override int getScreenSize(){
            return this.screenSize;
        }
        public override int getCameraPixels(){
            return this.cameraPixels;
        }
    }
}