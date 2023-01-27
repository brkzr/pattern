namespace   AdapterPattern{
    public interface ICard { 
        string GetCardDetail(int customerNumber);
        int GetTotalAmount(int cardNumber);
    }
}