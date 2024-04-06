import React, { useContext, useEffect, useCallback } from 'react'
import './style.css'

import { useParams, useNavigate } from "react-router-dom";
import { Link } from 'react-router-dom'
import { Segment, Message, Header, Image, Grid, Button, Placeholder, SegmentGroup, PlaceholderHeader, Table } from 'semantic-ui-react'

import API from 'Api'
import { useRequest } from 'Shared/Hooks'
import { Post } from 'Shared/Models'
import SimplePage from 'Shared/SimplePage'
import { UserContainer } from 'Shared/UserContainer'

const Posts = () => {
  const [loading, error, run, posts] = useRequest([])
  const { user } = useContext(UserContainer)

  useEffect(() => {
    run(API.getPosts())
  }, [run])

  return (
    <SimplePage icon='copy outline' title='Products' error={error}>
      <Grid columns={2}>
        <Grid.Column floated='left'>
          <p>This page fetches some protected data that only the logged in user ({user.email}) can see!</p>
        </Grid.Column>
        <Grid.Column textAlign='right'>
          <Button as={Link} to='/post/create' primary icon='plus' content='Add Product' />
        </Grid.Column>
      </Grid>

      {loading && <PostsPlaceholder />}
      {posts.length === 0 && !loading &&
        <Message warning>No posts found...</Message>}
      {<PostsTable posts={posts} />}
    </SimplePage>
  )
}

export default Posts;

const PostsPlaceholder = () => (
  <SegmentGroup style={{ marginBottom: '1em' }}>
    <Segment attached='bottom'>
      <Placeholder>
        <Placeholder.Header>
          <Placeholder.Line />
        </Placeholder.Header>
      </Placeholder>
    </Segment>
    <Segment attached='top'>
      <Placeholder>
        <Placeholder.Paragraph>
          <Placeholder.Line />
          <Placeholder.Line />
        </Placeholder.Paragraph>
      </Placeholder>
    </Segment>
  </SegmentGroup>
)

// const SinglePost = ({ id, title, body }: Post) => (
//   <Segment.Group key={id}>
//     <Header attached='top' as='h3'>
//       <Link to={`/post/${id}`}>{title}</Link>
//     </Header>
//     <Segment attached='bottom' content={body} />
//   </Segment.Group>
// )


const PostsTable = ({ posts }: any) => (
  <Table celled>
    <Table.Header>
      <tr>
        <Table.HeaderCell> </Table.HeaderCell>
        <Table.HeaderCell> Product Name</Table.HeaderCell>
        <Table.HeaderCell> SKU</Table.HeaderCell>
        <Table.HeaderCell> Variant</Table.HeaderCell>
        <Table.HeaderCell> Price</Table.HeaderCell>
        <Table.HeaderCell> Status</Table.HeaderCell>
        <Table.HeaderCell> Edit/Delete</Table.HeaderCell>
      </tr>
    </Table.Header>
    <tbody>
      {posts.map((post: any) => (
        <SinglePost key={post.id} {...post} />
      ))}
    </tbody>
  </Table>
);

const SinglePost = ({ id, title, body }: any) => {
  const [loading, error, run] = useRequest({} as Post);

  const handleDelete = useCallback(() => {
    run(API.deletePost(id), () => {
      window.location.reload(); // Reload the page after deletion
    });
  }, [run, id]);

  const text = "Ws38790";
  const randomNumber = generateRandomNumber(text);

  return (
    <Table.Row>
      <Table.Cell>
        <Image src="https://img.freepik.com/free-psd/wine-bottle-isolated-transparent-background_191095-25809.jpg" alt="Wine Bottle" size="tiny" />
      </Table.Cell>
      <Table.Cell>{title}</Table.Cell>
      <Table.Cell>{randomNumber}</Table.Cell>
      <Table.Cell>{body}</Table.Cell>
      <Table.Cell>$59</Table.Cell>
      <Table.Cell>Draft</Table.Cell>
      <Table.Cell>
        <div style={{ display: 'flex', gap: '8px' }}>
          <Button icon="edit" color="blue" as={Link} to={`/post/${id}/edit`} />
          <Button icon="delete" color="red" onClick={handleDelete} />
        </div>
      </Table.Cell>
    </Table.Row>
  );
};


// TODO: Move to utils
function generateRandomNumber(text: string) {
  // Extract digits from the text using regular expression
  const digits = text.match(/\d/g);

  // If digits are found in the text
  if (digits !== null) {
    // Concatenate all digits into a single string
    const digitString = digits.join('');

    // Convert the digit string to a number
    const number = parseInt(digitString, 10);

    // Generate a random number within the range of the extracted number
    const randomNumber = Math.floor(Math.random() * number) + 1;

    return randomNumber;
  } else {
    return null; // Return null if no digits are found in the text
  }
}