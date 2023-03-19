
namespace creational.factory_pattern{

    public class IPhonePlus : IPhone{
        
        public IPhonePlus(StorageOptions storage ){
            this.storage = storage;
            this.screenSize = 5;
            this.cameraPixels = 12;
            this.phoneType = PhoneType.Plus;
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