<?
   include (‘ds_fns.php’);
   session_start();

   // check if we have created our session variable
   if(!session_is_registered(‘expanded’))
   {
      $expanded = array();
      session_register(‘expanded’);
    }

    // check if an expand button was pressed
    // expand might equal ‘all’ or a postid or not be set
    if($expand)
    {
       if($expand == ‘all’)
          expand_all($expanded);
        else
           $expanded[$expand] = true;
    }

    // check if a collapse button was pressed
    // collapse might equal all or a postid or not be set
    if($collapse)
    {
      if($collapse==”all”)
      unset($expanded);
    else
      unset($expanded[$collapse]);
    }
    do_html_header(“Threads”);

    display_index_toolbar();

    // display the tree view of conversations
    display_tree($expanded);

    do_html_footer();
?>