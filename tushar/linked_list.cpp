#include<iostream>
using namespace std;

typedef struct Data{
    string Username;
}Data;

typedef struct node{
    Data* data;
    struct node * next;
}node;

node * start;
node * poo;

node* createNode(string data) {
    node* newNode = new node();
    auto dataObj = new Data();
    dataObj->Username= data;
    newNode->data = dataObj;
    newNode->next = NULL;
    return newNode;
}

node * getLastNode(node * head){
    node * p = head; 

    while(p->next!=NULL){
        p=p->next;
    }

    return p;
}

void printLL(node *head){
    node * p = head;

    while(p!=NULL){
        cout<<p->data->Username<<" ";
        p=p->next;
    }

}

int main(){
    start = NULL;
    poo=NULL;
    int nodes;
    cin>>nodes;
    while(nodes!=0){
        string data ;
        cin>>data;

        auto newNode = createNode(data);

        if(start==NULL){
            start = newNode;
            poo=start;

        } else {
            //auto lastNode = getLastNode(start);
            //lastNode->next = newNode;
            poo->next=newNode;
            poo=poo->next;

        }

        printLL(start);
        nodes--;
    }

    return 0;
}