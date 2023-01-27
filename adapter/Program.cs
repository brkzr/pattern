namespace AdapterPattern 
{
    internal class Program
    {
        static void Main(string[] args)
        {
            var cards = new Type[] { typeof(CreditCard), typeof(BankCard), typeof(PrePaidCardAdapter) };
            PrePaidCard prePaidCard = new PrePaidCard();
            foreach (var cardType in cards)
            {
                var cardObject = (cardType.Equals(typeof(PrePaidCardAdapter)) ? Activator.CreateInstance(cardType, prePaidCard) : Activator.CreateInstance(cardType)) as ICard;
                Console.WriteLine(cardObject?.GetCardDetail(123));
                Console.WriteLine($"Total Amount for {cardObject?.GetType().Name} : {cardObject?.GetTotalAmount(1234)}");
                Console.WriteLine("----------------------------");
            }


            // // Create a new instance of the BankCard class.
            // ICard card = new BankCard();
            // Console.WriteLine(card.GetCardDetail(123));
            // Console.WriteLine(card.GetTotalAmount(1234));

            // // Create a new instance of the CreditCard class.
            // card = new CreditCard();
            // Console.WriteLine(card.GetCardDetail(123));
            // card.GetTotalAmount(1234);

            // // Create a new instance of the PrePaidCard class.
            // PrePaidCard prePaidCard = new PrePaidCard();
            // card = new PrePaidCardAdapter(prePaidCard);
            // card.GetCardDetail(123);
            // card.GetTotalAmount(1234);
        }
    }
}