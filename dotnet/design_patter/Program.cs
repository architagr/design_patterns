using creational.factory_pattern;

namespace designPattern
{

    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("**** India Store ****");
            Store indiaStore = new IndiaStore();
            Console.WriteLine("created order for pro 128 GB");

            IPhone phone = indiaStore.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
            
            Console.WriteLine($"Screen Size - {phone.getScreenSize()}");
            Console.WriteLine($"Camera - {phone.getCameraPixels()} Pixels");
            phone.takePic();
            phone.toggleCameraShutterSound();
            phone.takePic();
            SimType[] sims = phone.getSimTypes();
            Console.WriteLine($"sim types {sims[0]}, {sims[1]}");


            Console.WriteLine("**** Japan Store ****");
            Store japanStore = new JapanStore();
            Console.WriteLine("created order for pro 128 GB");

            phone = japanStore.CreateOrder(PhoneType.Pro, StorageOptions.Gb128);
            
            Console.WriteLine($"Screen Size - {phone.getScreenSize()}");
            Console.WriteLine($"Camera - {phone.getCameraPixels()} Pixels");
            phone.takePic();
            phone.toggleCameraShutterSound();
            phone.takePic();
            
            sims = phone.getSimTypes();
            Console.WriteLine($"sim types {sims[0]}, {sims[1]}");

        }
    }
}