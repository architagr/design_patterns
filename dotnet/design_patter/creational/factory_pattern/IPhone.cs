namespace creational.factory_pattern{
    public abstract class IPhone  {
        protected int screenSize, cameraPixels;
	    protected PhoneType phoneType;
	    protected StorageOptions storage;
        protected bool canTurnOffShutterSound, shutterSound;
        protected SimType[] sims;

        public bool InstallSoftwares(){
            Console.WriteLine($"please wait we are updating the softwares for your {this.phoneType}, {this.storage}.");
            Console.WriteLine($"All softwares updated for {this.phoneType}, {this.storage}.");
            return true;
        }

        public bool Package(){
            Console.WriteLine($"{this.phoneType}, {this.storage} is getting packed !!!");
            Console.WriteLine($"{this.phoneType}, {this.storage} is packed !!!, PARTY");
            return true;
        }
        public PhoneType getPhoneType(){
            return this.phoneType;
        }
        public StorageOptions getStorage(){
            return this.storage;
        }
        public int getScreenSize(){
            return this.screenSize;
        }
        public int getCameraPixels(){
            return this.cameraPixels;
        }
        public void toggleCameraShutterSound(){
            if (this.canTurnOffShutterSound){
                this.shutterSound = !this.shutterSound;
            }
        }
        public SimType[] getSimTypes() {
            return this.sims;
        }
        public void takePic(){
            Console.WriteLine($"Shutter sound id {this.shutterSound}");   
        }
    }
}