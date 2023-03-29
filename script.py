import requests

BASE_URL = 'http://localhost:8000/'

def send_request(method, endpoint, data=None):
    url = BASE_URL + endpoint
    try:
        response = getattr(requests, method.lower())(url, json=data)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.HTTPError as errh:
        raise Exception(f"An HTTP Error occurred: {errh}")
    except requests.exceptions.ConnectionError as errc:
        raise Exception(f"An Error Connecting to the API occurred: {errc}")
    except requests.exceptions.Timeout as errt:
        raise Exception(f"A Timeout Error occurred: {errt}")
    except requests.exceptions.RequestException as err:
        raise Exception(f"An Unknown Error occurred: {err}")

if __name__ == '__main__':
    method = input('Select HTTP method (GET, POST, PUT, DELETE): ').upper()
    endpoint = input('Enter endpoint: ')

    if method in ['POST', 'PUT']:
        data = {}
        while True:
            key = input('Enter data key (leave empty to finish): ')
            if not key:
                break
            value = input(f'Enter value for "{key}": ')
            data[key] = value
    else:
        data = None

    try:
        response = send_request(method, endpoint, data)
        if method == 'DELETE':
            id_ = response.get('id')
            print(f'Deleted object with ID: {id_}')
        else:
            print(response)
    except Exception as e:
        print(e)