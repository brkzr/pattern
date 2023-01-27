namespace  AdapterPattern{
     public class BankCard : ICard
    {
         public string GetCardDetail(int customerNumber)
        {
             return  $"Bank Card Detail for {customerNumber}";
        }

         public int GetTotalAmount(int cardNumber)
        {
             return  100;
        }
    }
} 