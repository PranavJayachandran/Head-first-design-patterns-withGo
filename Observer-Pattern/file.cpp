#include <bits/stdc++.h>
using namespace std;


//The observer subscribes the subject and gets notified whenever the value in the subject is updated.
//Its important to notice that the observers and subject are always working on the interface rather than the implementation
//Look how Files and ReverseFiles are subscribing to the same object and how the subject doesn't care the types for those classes,
// rather its just important that the interfaces are implemented.

int dataCounter = 1;

class Subject;
class Observer;
class Files;
class DataSource;
class Observer
{
public:
    virtual void subscribe() = 0;
    virtual void unsubscribe() = 0;
    virtual void update(string s) = 0;
    virtual string getFileName() = 0 ;
    virtual ~Observer() = default;
};
class Subject
{
public:
    virtual void generateData() = 0;
    virtual void notifyObservers() = 0;
    virtual void addObserver(Observer *obj) = 0;
    virtual void removeObserver(Observer *obj) = 0;
    virtual ~Subject() = default;
};
class Files : public Observer
{
    Subject *dataSource;
    string fileName;

public:
    Files(Subject *_dataSource, string _fileName)
    {
        dataSource = _dataSource;
        fileName = _fileName;
    }
    void subscribe() override
    {
        dataSource->addObserver(this);
    }
    void unsubscribe() override
    {
        cout << "Unsubscribing" << fileName << "\n";
        dataSource->removeObserver(this);
    }
    void update(string data) override
    {
        ofstream file;
        file.open(fileName, ios::app);
        file << data;
        file.close();
    }
    string getFileName() override
    {
        return fileName;
    }
};
class ReverseFiles : public Observer
{
    Subject *dataSource;
    string fileName;

public:
    ReverseFiles(Subject *_dataSource, string _fileName)
    {
        dataSource = _dataSource;
        fileName = _fileName;
    }

    void subscribe() override
    {
        dataSource->addObserver(this);
    }
    void unsubscribe() override
    {
        cout << "Unsubscribing" << fileName << "\n";
        dataSource->removeObserver(this);
    }
    void update(string data) override
    {
        ofstream file;
        file.open(fileName, ios::app);
        reverse(data.begin(),data.end());
        file << data;
        file.close();
    }
    string getFileName() override
    {
        return fileName;
    }
};
class DataSource : public Subject
{
    vector<Observer *> fileObservers;
    string newData = "";

public:
    void generateData() override
    {
        newData = "Data" + to_string(dataCounter++) + "\n";
    }
    void notifyObservers() override
    {
        for (auto obj : fileObservers)
        {
            obj->update(newData);
        }
    }
    void addObserver(Observer *obj) override
    {
        fileObservers.push_back(obj);
    }
    void removeObserver(Observer *obj) override
    {
        fileObservers.erase(find(fileObservers.begin(), fileObservers.end(), obj));
    }
};
int main()
{
    DataSource *source = new DataSource();
    vector<Observer *> observers;
    while (1)
    {
        cout << "List of subscribed files\n";
        for (auto p : observers)
        {
            cout << p->getFileName() << "\n";
        }
        cout << "\n";

        cout << "Choose an option\n1. Push new data\n2. Add a new observer\n3. Remove an observer\n";
        int option;
        cin >> option;
        switch (option)
        {
        case 1:
        {
            source->generateData();
            source->notifyObservers();
            break;
        }
        case 2:
        {
            cout << "Enter a fileName\n";
            string fileName;
            cin >> fileName;
            cout<<"Should it be reversed or not(1 or 0)";
            int isReversed;
            cin>>isReversed;
            Observer *file;
            if(!isReversed)
                file = new Files(source, fileName);
            else
                file = new ReverseFiles(source,fileName);
            file->subscribe();
            observers.push_back(file);
            break;
        }
        case 3:
        {
            cout << "Enter the fileName of the observer to remove\n";
            string fileName;
            cin >> fileName;
            Observer *temp = NULL;
            for (auto p : observers)
            {
                if (p->getFileName() == fileName)
                {
                    temp = p;
                    break;
                }
            }
            if (temp != NULL)
            {
                temp->unsubscribe();
                observers.erase(find(observers.begin(), observers.end(), temp));
            }
            break;
        }
        default:
        {
            cout << "enter a correct value\n";
        }
        }
    }
}
