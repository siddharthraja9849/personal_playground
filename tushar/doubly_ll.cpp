#include<iostream>
using namespace std;

typedef struct node {
    int data;
    struct node * next;
    struct node * prev;
}node;

node* createNode(int data){
    node* newnode= new node();//initialise new node
    newnode->data=data;
    newnode->next=NULL;
    newnode->prev=NULL;
    return newnode;

}

node* getLastNode(node * start){
    node *p=start;
    while(p->next!=NULL){
        p=p->next;
    }
    return p;
}

void printdll(node * start){
    node *p = start;
    while(p!=NULL){
        cout<<p-> data <<"  ";
        p=p->next;
    }
}
void printreverse(node *end){
    node * q=end;
    while(q!=NULL){
        cout<< q -> data<<" ";
        q=q->prev;

    }
}

node * start; 
node * endd;

int main(){
    int n;
    cin>>n;
    start=NULL;
    endd=NULL;
    for(int i=1;i<=n;i++){
        int val;
        
        
        cin>>val;
        node * newnode = createNode(val);

        if(start==NULL){
            start=newnode;
            endd=newnode;
        }
        else{
            //auto lastNode=getLastNode(start);
            //lastNode->next=newnode;
            //newnode->prev=lastNode;
            endd->next=newnode;
            newnode->prev=endd;
            endd=endd->next;

        }
    }
    printdll(start);
    cout<<endl;
    printreverse(endd);

    return 0;
}