dactyl
======

``dactyl`` is a small and simple demo web service.  It exists only so
that `ida <https://github.com/cleardataeng/demo-container-ida>`_ can
provide a simple example of composing services.

``dactyl`` supports a single endpoint at ``/``, and only handles ``GET``
requests.

A ``GET /`` response is a JSON data structure with the following
elements:

* ``Hostname``: the hostname where ``dactyl`` is running
* ``RemoteAddr``: the IP address from which ``dactyl`` received your
  request
