namespace AdapterPattern
{
    public class PrePaidCardAdapter : ICard
    {
        private readonly PrePaidCard _prePaidCard;
        public PrePaidCardAdapter(PrePaidCard prePaidCard)
        {
            _prePaidCard = prePaidCard;
        }
        public string GetCardDetail(int customerNumber)
        {
            return _prePaidCard.GetPrepaidDetail(customerNumber.ToString());
        }

        public int GetTotalAmount(int cardNumber)
        {
            return _prePaidCard.GetAmount(cardNumber.ToString());
        }
    }
}