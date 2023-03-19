
namespace creational.factory_pattern{
    public class Store {

        public IPhone CreateOrder(PhoneType type, StorageOptions storage){
            IPhone phone = null;
            if (type == PhoneType.Pro){
                phone = new IPhonePro(storage);
            } else if (type == PhoneType.Plus) {
                phone = new IPhonePlus(storage);
            }

            phone.InstallSoftwares();
            phone.Package();
            return phone;
        }
    }

}