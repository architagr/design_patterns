
namespace creational.factory_pattern{

    public class IPhonePlus : IPhone{
        
        public IPhonePlus(StorageOptions storage ){
            this.storage = storage;
            this.screenSize = 5;
            this.cameraPixels = 12;
            this.phoneType = PhoneType.Plus;
        }
    }
}