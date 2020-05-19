package main;

type StockSpanner struct {
    prices []int
}


func Constructor() StockSpanner {
	return StockSpanner{ prices: []int {}}
}


func (this *StockSpanner) Next(price int) int {

	span := 1;
	for i:= len(this.prices) -1; i >= 0; i--{
		if price >= this.prices[i]{
			span++;
		}else{
			break;
		}
	}

	this.prices = append(this.prices, price);

    return span;
}

func main(){
	obj := Constructor();
	
	result := obj.Next(31);
	result = obj.Next(41);
	result = obj.Next(48);
	result = obj.Next(59);
	result = obj.Next(79);
	print(result);
}