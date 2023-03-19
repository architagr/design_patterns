
namespace creational.factory_pattern{

    public class IPhonePlus : IPhone{
        
        public IPhonePlus(StorageOptions storage, bool canTurnOffShutterSound, SimType[] simTypes ){
            this.storage = storage;
            this.screenSize = 5;
            this.cameraPixels = 12;
            this.phoneType = PhoneType.Plus;
            this.canTurnOffShutterSound = canTurnOffShutterSound;
            this.shutterSound = true;
            this.sims = simTypes;
        }
    }
}