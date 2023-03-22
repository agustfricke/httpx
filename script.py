import requests

usr_req = input('''
                Select the http method 
                # 1 - GET
                # 2 - POST 
                # 3 - PUT / {id}
                # 4 - DELETE / {id}
                = ''')
base_url = input('''
            
            Select base URL:
            # 1 - http:/127.0.0.1:8000/
            # 2 - http:localhost:3500/
            = ''')

miss_url = input('Set the missing param: ')

if base_url == '1':
    url = 'http://127.0.0.1:8000/' + miss_url + '/'
elif base_url == '2':
    url = 'http://localhost:3500/' + miss_url + '/'


if usr_req == '1':
    # url = input('set url: ')
    r = requests.get(url)
    print(r.json()) 
elif usr_req == '2':
    # url = input('set url: ')
    key = input('key: ')
    value = input('value: ')
    r = requests.post(url, data={key: value})
    print(r.json()) 
elif usr_req == '3':
    # url = input('set url: ')
    key = input('key: ')
    value = input('value: ')
    r = requests.put(url, data={key: value})
    print(r.json()) 
elif usr_req == '4':
    # url = input('set url: ')
    r = requests.delete(url)
    print(r) 
else:
    print('Dont know! :(')



