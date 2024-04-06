import React, { useContext, useEffect,useCallback } from 'react'
import './style.css'

import { useParams, useNavigate } from "react-router-dom";
import { Link } from 'react-router-dom'
import { Segment, Message, Header,Grid, Button, Placeholder, SegmentGroup, PlaceholderHeader,Table} from 'semantic-ui-react'

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

      {loading && <PostsPlaceholder/>}
      {posts.length === 0 && !loading && 
        <Message warning>No posts found...</Message>}
      {<PostsTable posts={posts}/>}
    </SimplePage>
  )
}

export default Posts;

const PostsPlaceholder = () => (
  <SegmentGroup style={{ marginBottom: '1em' }}>
    <Segment attached='bottom'>
      <Placeholder>
        <Placeholder.Header>
          <Placeholder.Line/>
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


const PostsTable = ({ posts }:any) => (
  <Table celled>
     <Table.Header>
      <tr>
        <Table.HeaderCell> Product</Table.HeaderCell>
        <Table.HeaderCell> Title</Table.HeaderCell>
        <Table.HeaderCell> Status</Table.HeaderCell>
        <Table.HeaderCell> Image</Table.HeaderCell>
        <Table.HeaderCell> Edit/Delete</Table.HeaderCell>
      </tr>
    </Table.Header>
    <tbody>
      {posts.map((post:any) => (
        <SinglePost key={post.id} {...post} />
      ))}
    </tbody>
  </Table>
);

const SinglePost = ({ id, title, body }: any) => {
  const [loading, error, run] = useRequest({} as Post)

  const handleDelete = useCallback(() => { // TODO: To be moved to parent context
    run(API.deletePost(id), () => {
      window.location.reload(); // TODO: Refetch API instead of making window.location
    })
  }, [run, id])

  return (
  <tr>
    <td>
      <Link to={`/post/${id}`}>{title}</Link>
    </td>
    <td>{body}</td>
    {/* Todo: Add status */}
    <td>Draft</td> 
    {/* Todo: Add status */}
    <td>{body}</td>
    <td>
      <Link to={`/post/${id}/edit`}><Button icon="edit" color="blue"/></Link>
      <Button icon="delete" color="blue" onClick={handleDelete}/>
    </td>

  </tr>
)};