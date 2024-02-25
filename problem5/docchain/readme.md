\`\`\`markdown
# DocChain - A Blockchain Application for Document Management

DocChain is a blockchain application built with Ignite CLI and Cosmos SDK. It allows users to create, read, update, and delete (CRUD) documents. Each document has a `title` and a `body`.

## Getting Started

1. **Install All Dependencies**: Install Go and Ignite CLI 

2. **Start the Blockchain Node**: Start your blockchain node locally using the `ignite chain serve` command:

```
ignite chain serve
```

## CRUD Operations

### Create a Document

To create a document, use the `create-document` command followed by the title and body of the document:

```
docchaind tx docchain create-document "My Document Title" "This is the body of my document."
```

### Read a Document

To read a document, use the `show-document` command followed by the ID of the document:

```
docchaind q docchain show-document [ID]
```

### Update a Document

To update a document, use the `update-document` command followed by the ID, new title, and new body of the document:

```
docchaind tx docchain update-document [ID] "New Document Title" "This is the new body of my document."
```

### Delete a Document

To delete a document, use the `delete-document` command followed by the ID of the document:

```
docchaind tx docchain delete-document [ID]
```

### Get All Documents

To get all documents, use the `list-document` command:

```
docchaind q docchain list-document
```

### Search Document by Title

To search for a document by title, use the `search-title` command followed by the word you want to search for in the document titles:

```
docchaind q docchain search-title [word]
```

## Conclusion

DocChain provides a simple and intuitive interface for managing documents on a blockchain. With its robust CRUD operations and additional functionalities, users can easily create, read, update, and delete documents, as well as retrieve all documents and search for documents by title. We hope you find DocChain useful and look forward to seeing what you build with it!

