namespace creational.factory_pattern{
    public abstract class Store {

        public IPhone CreateOrder(PhoneType type, StorageOptions storage){
            IPhone phone = this.createPhone(type, storage);
            phone.InstallSoftwares();
            phone.Package();
            return phone;
        }
        public abstract IPhone createPhone(PhoneType type, StorageOptions storage);
    }
    public class IndiaStore : Store{
        private SimType[] simTypes;
        private bool canTurnOffShutterSound;
        public IndiaStore(){
            this.canTurnOffShutterSound = true;
            this.simTypes = new SimType[]{SimType.Esim, SimType.PhysicalSim};
        }
        public override IPhone createPhone(PhoneType type, StorageOptions storage){
            IPhone phone = null;
            if (type == PhoneType.Pro){
                phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
            } else if (type == PhoneType.Plus) {
                phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
            }
            return phone;
        }
    }
    public class JapanStore : Store{
        private SimType[] simTypes;
        private bool canTurnOffShutterSound;
        public JapanStore(){
            this.canTurnOffShutterSound = false;
            this.simTypes = new SimType[]{SimType.PhysicalSim, SimType.PhysicalSim};
        }
        public override IPhone createPhone(PhoneType type, StorageOptions storage){
            IPhone phone = null;
            if (type == PhoneType.Pro){
                phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
            } else if (type == PhoneType.Plus) {
                phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
            }
            return phone;
        }
    }

    public class USAStore : Store{
        private SimType[] simTypes;
        private bool canTurnOffShutterSound;
        public USAStore(){
            this.canTurnOffShutterSound = true;
            this.simTypes = new SimType[]{SimType.Esim, SimType.Esim};
        }
        public override IPhone createPhone(PhoneType type, StorageOptions storage){
            IPhone phone = null;
            if (type == PhoneType.Pro){
                phone = new IPhonePro(storage, this.canTurnOffShutterSound, this.simTypes);
            } else if (type == PhoneType.Plus) {
                phone = new IPhonePlus(storage, this.canTurnOffShutterSound, this.simTypes);
            }
            return phone;
        }
    }
}