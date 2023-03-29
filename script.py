import requests

BASE_URL = 'http://localhost:8000/'

def send_request(method, endpoint, data=None):
    url = BASE_URL + endpoint
    try:
        response = getattr(requests, method.lower())(url, json=data)
        if response.content:  # check if response is not empty
            response.raise_for_status()
            return response.json()
        elif response.status_code == 204:
            # return success message with ID from the URL
            id_ = endpoint.split('/')[-2]
            return {'message': f'Deleted object with ID: {id_}'}
        else:
            raise Exception(f"An unknown error occurred. Response content: {response.content}")
    except requests.exceptions.HTTPError as errh:
        print(f"An HTTP Error occurred: {errh}")
    except requests.exceptions.ConnectionError as errc:
        print(f"An Error Connecting to the API occurred: {errc}")
    except requests.exceptions.Timeout as errt:
        print(f"A Timeout Error occurred: {errt}")
    except requests.exceptions.RequestException as err:
        print(f"An Unknown Error occurred: {err}")

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

    response = send_request(method, endpoint, data)
    if response:
        print(response)