# API Mocker

This project helps in mocking APIs for a given JSON structure. Even though lots of project exist to create a mock server
based on swagger definition, because of verbosity of swagger file, it ends up becoming complicated. This project aims to
have a very simple configuration to quickly run a mock server on local machine for development. 


## Mocking an API
Install the binary using: `go install github.com/vikash/api-mocker@latest`. Once installed, it can be executed by simply
typing `api-mocker` in terminal as long as `$GOPATH/bin` is added to `$PATH`. 

A json structure for a model needs to be written. Some samples are written in the `models` folder. By default, this
projects is reading the definitions from `./models` folder. Server starts on port 8000. 

Following methods are supported:

* GET /{entity} will give a list of entities.

## Data Types
Fields in entity definition file can be of several data types. 
Depending on the data type, the mocker provides appropriate values while mocking an API. The data types will 
also be used to validate parameters in POST, PUT or PATCH requests in future. 

It is strongly advised to use the most specific data format for any field. For example, 
even though image can be denoted as string with URL; one should always use image type in the context of Image 
as it provides for more accurate environment and proper mocking and testing can be done with more specific information.


### Predefined data types

#### number
This indicates a number. min and max values can be provided as additional keys. 
Default values are 1 and 999999999 respectively.

#### string
This is used for text values. Optional minLength and maxLength can be provided to restrict the size of the string. 
Default values are 4 and 20 respectively.

#### image
This is used to indicate an image. If an optional size is provided in {width}X{height} format, a URL for placeholder image of this size is return.
Instead of size, one can choose to provide one or more of minWidth, maxWidth and aspectRatio in format of 'W:H' like 16:9 or 4:3. The default values are 160, 1000 and 1:1 respectively. Unit of all sizes are pixels.

#### array
array
In order to use array, one has to provide item, which indicates the data type of the elements of the array. For example, if an entity can have multiple images, following definition can be provided:

```json
{
    "images": {
        "type": "array",
        "maxLength": 10,
        "minLength": 2,
        "item": {
            "type": "image",
            "aspect-ratio": "4:3"
        }
    }
}
```
Note that `item` can have all keys defined just like any normal field definition.

#### Custom Data Type
All the defined models can be used as a type as well to support nested behaviour. For example,
if we have a model in category.json, we can use this model as a type in another model like book.json as shown below:
```json
// book.json
{
  "id": {
    "type": "number",
    "min": "1",
    "required": true
  },
  "name": {
    "type": "string",
    "required": true
  },
  "imageURL": {
    "type": "image",
    "size": "160x160"
  },
  "category": {
    "type": "category" // NOTE THIS
  }
}
```