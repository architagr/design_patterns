
namespace creational.factory_pattern{
    public abstract class PhoneFactory {

        public IPhone createPhone(PhoneType type, StorageOptions storage){
            IPhone phone = null;
            if (type == PhoneType.Pro){
                phone = new IPhonePro(storage);
            } else if (type == PhoneType.Plus) {
                phone = new IPhonePlus(storage);
            }
            return phone;
        }
    }
}