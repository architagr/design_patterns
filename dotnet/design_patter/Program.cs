using creational.factory_pattern;

namespace designPattern
{

    class Program
    {
        static void Main(string[] args)
        {
            IPhone pro128 = new IPhonePro(StorageOptions.Gb128);
            Console.WriteLine("*** Pro ***");

            Console.WriteLine($"Screen size {pro128.getScreenSize()}");
            Console.WriteLine($"Screen size {pro128.getCameraPixels()}");

           IPhone plus128 = new IPhonePlus(StorageOptions.Gb128);
            Console.WriteLine("*** Plus ***");
            Console.WriteLine($"Screen size {plus128.getScreenSize()}");
            Console.WriteLine($"Screen size {plus128.getCameraPixels()}");

            Console.WriteLine("*** Pro from store ***");

            Store store = new Store();
            IPhone phone = store.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
            Console.WriteLine($"Screen size {phone.getScreenSize()}");
            Console.WriteLine($"Screen size {phone.getCameraPixels()}");
        }
    }
}