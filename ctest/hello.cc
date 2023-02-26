using namespace std; 

int main() 
{ 
	 
	string str = "";
	cout<<"Enter the string:\n";
	
	cin>>str;

	char arr[str.length() + 1]; 

	strcpy(arr, str.c_str()); 
    cout<<"String to char array conversion:\n";
	for (int i = 0; i < str.length(); i++) 
		cout << arr[i]; 

	return 0; 
} 
